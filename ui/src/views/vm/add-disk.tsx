import { useMutation, useQuery } from "@tanstack/react-query";
import {
    CreateResourceInput,
    Resource,
    ResourceKind,
} from "../../types/resource";
import { get } from "../../data/queries/get";
import { Machine, MachineDisk, MachineStatus } from "../../types/machine";
import { useParams } from "react-router";
import {
    Box,
    Button,
    Card,
    Container,
    Flex,
    Text,
    TextField,
} from "@radix-ui/themes";
import { Form, Formik, FormikHelpers } from "formik";
import { convert, Units } from "../../utils/conversion";
import { useNavigate } from "react-router";
import { PoolSelect } from "../../components/pool-select";
import { create } from "../../data/mutations/create";
import { patch } from "../../data/mutations/patch";
import { Volume } from "../../types/volume";
import * as jsonpatch from "fast-json-patch";

const generateUUID = (): string => {
    // @ts-ignore
    return ([1e7] + -1e3 + -4e3 + -8e3 + -1e11).replace(/[018]/g, (c) =>
        (
            c ^
            (crypto.getRandomValues(new Uint8Array(1))[0] & (15 >> (c / 4)))
        ).toString(16),
    );
};

interface AddDiskValues {
    pool_id: string;
    capacity: number;
}

export const AddDiskView: React.FC<unknown> = () => {
    const { id } = useParams<{ id: string }>();
    const navigate = useNavigate();

    const {
        data: vm,
        isError,
        isLoading,
    } = useQuery({
        queryKey: [ResourceKind.Machine, id],
        queryFn: () => get<Machine, MachineStatus>(id!, ResourceKind.Machine),
    });

    const { mutate: createMutation } = useMutation({
        mutationFn: create<any>,
    });

    const { mutate: patchMutation } = useMutation({
        mutationFn: patch,
    });

    if (isError) {
        return <div>error</div>;
    }

    if (isLoading) {
        return <div>loading..</div>;
    }
    return (
        <Box p="9">
            <Container size="1">
                <Text size="6">Create Storage Volume</Text>
                <Card mt="2">
                    <Box pt="3">
                        <Formik<AddDiskValues>
                            initialValues={{ pool_id: "", capacity: 12 }}
                            onSubmit={(
                                values,
                                { setSubmitting }: FormikHelpers<AddDiskValues>,
                            ) => {
                                const capacity = convert(
                                    values.capacity,
                                    Units.Gigabyte,
                                    Units.Bytes,
                                );

                                const volumeName = generateUUID();
                                const volumeId =
                                    values.pool_id! + "/" + volumeName;

                                const input: CreateResourceInput<Volume> = {
                                    id: volumeId,
                                    owner_ref: {
                                        id: values.pool_id!,
                                        kind: ResourceKind.StoragePool,
                                    },
                                    annotations: {
                                        "machinereq.id":
                                            vm!.annotations!["machinereq.id"],
                                        "machinereq.name": vm!.spec!.name,
                                    },
                                    kind: ResourceKind.StorageVolume,
                                    spec: {
                                        name: volumeName,
                                        pool: "",
                                        capacity: capacity,
                                        allocation: 0,
                                        key: "",
                                    },
                                };

                                createMutation(input, {
                                    onSuccess: () => {
                                        const disk: MachineDisk = {
                                            volume: volumeId,
                                            file: {},
                                            key: "",
                                            device: "",
                                            type: "",
                                            target: {
                                                dev: "",
                                            },
                                        };

                                        const observer = jsonpatch.observe<
                                            Resource<Machine, MachineStatus>
                                        >(vm!);
                                        vm!.spec?.disks.push(disk);
                                        const patches =
                                            jsonpatch.generate(observer);

                                        patchMutation(
                                            {
                                                id: vm!.id,
                                                kind: vm!.kind,
                                                patches,
                                            },
                                            {
                                                onSuccess: () => {
                                                    setSubmitting(false);
                                                    navigate(
                                                        `/${ResourceKind.Machine}/${vm!.id}?tab=devices`,
                                                    );
                                                },
                                            },
                                        );
                                    },
                                });
                            }}
                        >
                            {({ setFieldValue }) => (
                                <Form>
                                    <Box>
                                        <Box pb="1">
                                            <label htmlFor="name">Pool</label>
                                        </Box>
                                        <PoolSelect
                                            nodeId={vm!.owner!.id}
                                            onChange={(e) => {
                                                setFieldValue("pool_id", e);
                                            }}
                                        />
                                    </Box>
                                    <Box pt="3">
                                        <Box pb="1">
                                            <label htmlFor="capacity">
                                                Capacity
                                            </label>
                                        </Box>
                                        <TextField.Root
                                            placeholder="40 GB"
                                            id="capacity"
                                            name="capacity"
                                            type="number"
                                            onChange={(e) =>
                                                setFieldValue(
                                                    "capacity",
                                                    e.target.value,
                                                )
                                            }
                                        ></TextField.Root>
                                    </Box>

                                    <Flex gap="3" mt="4" justify="end">
                                        <Button
                                            variant="soft"
                                            color="red"
                                            type="button"
                                            onClick={() => {
                                                navigate(
                                                    `/${ResourceKind.Machine}/${vm!.id}?tab=devices`,
                                                );
                                            }}
                                        >
                                            Cancel
                                        </Button>
                                        <Button
                                            variant="soft"
                                            color="purple"
                                            type="submit"
                                        >
                                            Create
                                        </Button>
                                    </Flex>
                                </Form>
                            )}
                        </Formik>
                    </Box>
                </Card>
            </Container>
        </Box>
    );
};
