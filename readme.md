## readme（说明）

- `v20/`，用反射和 Unsafe，实现一个简单的 ORM 框架。主要处理 MySQL 的语句。
  - 主要实现：
    - SELECT
    - 元数据
    - 结果集
    - INSERT
    - 方言（处理 MySQL 和 SQLite3 的 On Conflict 语句的差异）
    - UPDATE
    - DELETE
  - 次要实现：
    - 中间件
