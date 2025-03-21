import { useQuery } from "@tanstack/react-query";
import { useNavigate } from "react-router";
import { list } from "../../data/queries/list";
import { ResourceKind } from "../../types/resource";
import { Table, Badge, Card } from "@radix-ui/themes";
import { Node } from "../../types/node";
import { getUptime } from "../../utils/uptime";

export const NodesTab: React.FC<unknown> = () => {
    const navigate = useNavigate();
    const nodes = useQuery({
        queryKey: [ResourceKind.Instance, "self", ResourceKind.Node],
        queryFn: () =>
            list<Node>(ResourceKind.Node, "self", ResourceKind.Instance),
    });

    if (nodes.error) {
        return <div>Error fetching nodes</div>;
    }

    return (
        <Card>
            <Table.Root>
                <Table.Header>
                    <Table.Row>
                        <Table.ColumnHeaderCell>Status</Table.ColumnHeaderCell>
                        <Table.ColumnHeaderCell>Name</Table.ColumnHeaderCell>
                        <Table.ColumnHeaderCell>
                            Machines
                        </Table.ColumnHeaderCell>
                        <Table.ColumnHeaderCell>Uptime</Table.ColumnHeaderCell>
                    </Table.Row>
                </Table.Header>

                <Table.Body>
                    {nodes.data?.list?.map((node) => {
                        const color = "green";

                        const uptime = getUptime(node.spec!.uptime!);

                        return (
                            <Table.Row
                                key={node.id}
                                onClick={() => {
                                    localStorage.removeItem(`${node.id}/tab`);

                                    navigate(
                                        `/${ResourceKind.Node}/${node.id}`,
                                    );
                                }}
                            >
                                <Table.RowHeaderCell>
                                    <Badge color={color as any}>
                                        Connected
                                    </Badge>
                                </Table.RowHeaderCell>
                                <Table.RowHeaderCell>
                                    {node.spec!.name}
                                </Table.RowHeaderCell>
                                <Table.Cell>
                                    <Badge color="green">
                                        {node.spec!.active_machines}
                                    </Badge>
                                    /
                                    <Badge color="purple">
                                        {node.spec!.total_machines}
                                    </Badge>
                                </Table.Cell>
                                <Table.Cell>
                                    <Badge color="amber">{uptime}</Badge>
                                </Table.Cell>
                            </Table.Row>
                        );
                    })}
                </Table.Body>
            </Table.Root>
        </Card>
    );
};
