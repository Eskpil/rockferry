import { Button, Table, Text } from "@radix-ui/themes";
import { Machine, MachineDisk, MachineStatus } from "../../../types/machine";
import { Resource, ResourceKind } from "../../../types/resource";
import { useMutation, useQuery } from "@tanstack/react-query";
import { Volume } from "../../../types/volume";
import { get } from "../../../data/queries/get";
import { WithOwner } from "../../../components/withowner";
import { Pool } from "../../../types/pool";
import { convert, Units } from "../../../utils/conversion";
import { Badge } from "@radix-ui/themes/src/index.js";
import { checkVolumeAllocationStatus } from "../../../utils/allocationstatus";
import { useNavigate } from "react-router";
import { patch } from "../../../data/mutations/patch";
import * as jsonpatch from "fast-json-patch";

interface Props {
    vm: Resource<Machine, MachineStatus>;
}

interface DiskViewProps {
    vm: Resource<Machine, MachineStatus>;
    disk: MachineDisk;
}

const DiskView: React.FC<DiskViewProps> = ({ disk, vm }) => {
    const navigate = useNavigate();

    const {
        data: volume,
        isError,
        isLoading,
    } = useQuery({
        queryKey: [ResourceKind.StorageVolume, disk.volume],
        queryFn: () => get<Volume>(disk.volume, ResourceKind.StorageVolume),
    });

    const { mutate: patchMutation } = useMutation({
        mutationFn: patch,
    });

    if (isError) {
        return <div>error...</div>;
    }

    if (isLoading) {
        return <div>loading..</div>;
    }

    const capacity = Math.round(
        convert(volume!.spec!.capacity, Units.Bytes, Units.Gigabyte),
    );
    const allocation = Math.round(
        convert(volume!.spec!.allocation, Units.Bytes, Units.Gigabyte),
    );
    return (
        <WithOwner<Pool> res={volume!}>
            {({ owner }) => (
                <Table.Row>
                    <Table.RowHeaderCell>
                        <Text
                            className="hover:cursor-pointer"
                            onClick={() =>
                                navigate(
                                    `/${ResourceKind.StoragePool}/${owner.id}`,
                                )
                            }
                        >
                            {disk.volume.split("/")[1]}
                        </Text>
                    </Table.RowHeaderCell>
                    <Table.RowHeaderCell>
                        {volume?.spec?.type == "iso" ? (
                            <Badge color="amber">n/a</Badge>
                        ) : (
                            <>
                                <Badge
                                    color={
                                        checkVolumeAllocationStatus(
                                            capacity,
                                            allocation,
                                        ) as any
                                    }
                                >
                                    {allocation} gb
                                </Badge>
                                {" / "}
                                <Badge>{capacity} gb</Badge>
                            </>
                        )}
                    </Table.RowHeaderCell>
                    <Table.RowHeaderCell>{disk.target.dev}</Table.RowHeaderCell>
                    <Table.RowHeaderCell>
                        <Text
                            className="hover:cursor-pointer"
                            color="purple"
                            onClick={() =>
                                navigate(
                                    `/${ResourceKind.StoragePool}/${owner.id}`,
                                )
                            }
                        >
                            {owner.spec?.name}
                        </Text>
                    </Table.RowHeaderCell>
                    <Table.Cell>{disk.device}</Table.Cell>
                    <Table.Cell>{volume?.spec?.type}</Table.Cell>
                    <Table.Cell>
                        <Button
                            variant="soft"
                            color="red"
                            size="1"
                            onClick={() => {
                                const observer =
                                    jsonpatch.observe<
                                        Resource<Machine, MachineStatus>
                                    >(vm);

                                const index = vm.spec?.disks.indexOf(disk);
                                if (-1 >= index!) {
                                    console.log("some error here");
                                    return;
                                }

                                vm.spec?.disks.splice(index!, 1);

                                const patches = jsonpatch.generate(observer);

                                patchMutation({
                                    id: vm.id,
                                    kind: vm.kind,
                                    patches,
                                });
                            }}
                            style={{ width: "70%" }}
                        >
                            Remove
                        </Button>
                    </Table.Cell>
                </Table.Row>
            )}
        </WithOwner>
    );
};

export const DisksView: React.FC<Props> = ({ vm }) => {
    const navigate = useNavigate();

    return (
        <Table.Root>
            <Table.Header>
                <Table.Row>
                    <Table.ColumnHeaderCell>Volume</Table.ColumnHeaderCell>
                    <Table.ColumnHeaderCell>Size</Table.ColumnHeaderCell>
                    <Table.ColumnHeaderCell>Target</Table.ColumnHeaderCell>
                    <Table.ColumnHeaderCell>Pool</Table.ColumnHeaderCell>
                    <Table.ColumnHeaderCell>Device</Table.ColumnHeaderCell>
                    <Table.ColumnHeaderCell>Type</Table.ColumnHeaderCell>
                    <Table.ColumnHeaderCell>
                        <Button
                            variant="soft"
                            size="1"
                            style={{ width: "70%" }}
                            onClick={() => {
                                navigate(
                                    `/${ResourceKind.Machine}/${vm.id}/add-disk`,
                                );
                            }}
                        >
                            Add
                        </Button>
                    </Table.ColumnHeaderCell>
                </Table.Row>
            </Table.Header>

            <Table.Body>
                {vm.spec?.disks.map((disk) => (
                    <>
                        {disk.volume != "" ? (
                            <DiskView key={disk.volume} disk={disk} vm={vm} />
                        ) : undefined}
                    </>
                ))}
            </Table.Body>
        </Table.Root>
    );
};
