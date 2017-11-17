# 基于角色-资源的权限控制模型

用户属于什么角色 -- 可以一对多
什么角色可以对什么资源做什么操作

```sql
create table users(
    id bigserial not null primary key,
    name text not null,
    created_at timestamptz not null default clock_timestamp()
);

create table roles(
    id bigserial not null primary key,
    name text not null
);

create table entitys(
    id bigserial not null primary key,
    name text not null
);

create table ops(
    id bigserial not null primary key,
    name text not null
);

create table user_roles(
    user_id bigint,
    role_id bigint
);

create table role_entitys(
    role_id bigint,
    entity_id bigint
);

create table role_ops(
    role_id bigint,
    op_id bigint
);
```

资源是一直变化的（增删），增加的资源怎么判断什么角色可以使用呢，增加资源的时候指定？如果是由前端添加的呢，怎么指定？
如果修改了规则怎么办呢？
接口可以被什么角色的用户调用。
先判断用户所属角色是否可以调用该接口
再判断用户是否对资源拥有操作的权限。