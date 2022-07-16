-- 코드를 입력하세요
SET @hour = -1;
SELECT
    (@hour := @hour +1) as HOUR,
    (SELECT COUNT(DISTINCT ANIMAL_ID) FROM ANIMAL_OUTS WHERE HOUR(DATETIME) = @hour) as `COUNT`
from ANIMAL_OUTS
where @hour < 23