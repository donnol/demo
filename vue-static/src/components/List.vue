<template>
    <div class="list">
        <h1>1124</h1>
        <ul>
            <li>{{msg}}</li><br>
            <li>{{content}}</li><br>
        </ul>
        <h2><a>{{author}}</a></h2>
    </div>
</template>

<script>
export default {
  name: 'List',
  data() {
    return {
      msg: '以往寒风躲闪行，今秋尽是路人多。',
      content: `分布式系统之微服务架构设计要点
        1, 调度/业务流程管理: Kubernetes
        2, 服务发现和注册: etcd
        3, 系统状态管理/集群管理: Kubernetes Nodes/Masters
        4, 数据存储: Kubernetes Persistent volumes
        5, 网络: 
        6, 监测/审核/日志:
      `,
      author: 'jd'
    }
  }
}
</script>

<style scoped>
h1 {
  font-weight: normal;
  text-align: center;
  color: #42b983;
}
ul {
  list-style-type: none;
  padding: 0;
  text-align: center;
}
li {
  display: inline-block;
  margin: 0 10px;
}
h2 {
  color: #443322;
  text-align: right;
}
</style>