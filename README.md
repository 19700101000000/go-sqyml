# sqyml
Purpose : from yml to sql

## exqmple
```yml 
main:
  type: select
  table: foo
  column: bar
```
```sql
SELECT `bar` FROM `foo`;
```
