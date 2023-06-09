-- Find all persons associated with that organization

WITH RECURSIVE traverse(node_id, entity_type, entity_id) AS (
    SELECT
        node_id,
        entity_type,
        entity_id
    FROM
        entity
    WHERE
        entity.entity_id = '1' AND
        entity.entity_type = 'Person'
    UNION ALL
    SELECT
        entity.node_id,
        entity.entity_type,
        entity.entity_id
    FROM traverse JOIN
    relation ON traverse.node_id = relation.child JOIN
    entity ON relation.parent = entity.node_id
)
SELECT
    traverse.entity_id as order_id
FROM traverse
WHERE traverse.entity_type = 'Organization'
GROUP BY traverse.entity_id
ORDER BY traverse.entity_id ASC
