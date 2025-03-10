import { useQuery } from "@tanstack/react-query";
import { list } from "../../data/queries/list";
import { Badge, Card, Table } from "@radix-ui/themes";
import { useNavigate } from "react-router";
import { convert, Units } from "../../utils/conversion";
import { ResourceKind } from "../../types/resource";
import { Pool } from "../../types/pool";

interface Props {
    nodeId: string;
}

// TODO: Add skeleton in table body for clean ui when loading

export const PoolsView: React.FC<Props> = ({ nodeId }) => {
    const navigate = useNavigate();

    const data = useQuery({
        queryKey: [nodeId, `pools`],
        queryFn: () =>
            list<Pool>(ResourceKind.StoragePool, nodeId, ResourceKind.Node),
    });

    if (data.isError) {
        console.log(data.error);
        return <p>error</p>;
    }

    return (
        <Card>
            <Table.Root layout="auto">
                <Table.Header>
                    <Table.Row>
                        <Table.ColumnHeaderCell>Name</Table.ColumnHeaderCell>
                        <Table.ColumnHeaderCell>Default</Table.ColumnHeaderCell>
                        <Table.ColumnHeaderCell>Usage</Table.ColumnHeaderCell>
                        <Table.ColumnHeaderCell>Backend</Table.ColumnHeaderCell>
                    </Table.Row>
                </Table.Header>

                <Table.Body>
                    {data.data?.list.map((resource) => {
                        const pool = resource.spec!;

                        let default_ = "no";

                        if (
                            resource.annotations &&
                            resource.annotations["rockferry.default"] == "yes"
                        ) {
                            default_ = "yes";
                        }

                        const capacity_gb = Math.round(
                            convert(pool.capacity, Units.Bytes, Units.Gigabyte),
                        );
                        const allocated_gb = Math.round(
                            convert(
                                pool.allocation,
                                Units.Bytes,
                                Units.Gigabyte,
                            ),
                        );

                        return (
                            <Table.Row
                                key={pool.id}
                                onClick={() => {
                                    navigate(`/pools/${resource.id}`);
                                }}
                            >
                                <Table.RowHeaderCell>
                                    {pool.name}
                                </Table.RowHeaderCell>
                                <Table.Cell>{default_}</Table.Cell>
                                <Table.Cell>
                                    <Badge color="green">
                                        {allocated_gb} Gb
                                    </Badge>
                                    /
                                    <Badge color="purple">
                                        {capacity_gb} Gb
                                    </Badge>
                                </Table.Cell>
                                <Table.Cell>
                                    <Badge color="amber">{pool.type}</Badge>
                                </Table.Cell>
                            </Table.Row>
                        );
                    })}
                </Table.Body>
            </Table.Root>
        </Card>
    );
};
