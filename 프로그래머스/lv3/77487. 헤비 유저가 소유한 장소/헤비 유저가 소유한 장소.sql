-- 코드를 입력하세요
SELECT 
     PLACES.ID
    ,PLACES.NAME
    ,PLACES.HOST_ID
FROM PLACES
INNER JOIN (
    SELECT
         COUNT(HOST_ID) AS count_host_id
        ,HOST_ID
    FROM PLACES
    GROUP BY PLACES.HOST_ID
    HAVING count_host_id>=2
)AS place ON place.HOST_ID=PLACES.HOST_ID
ORDER BY PLACES.ID ASC