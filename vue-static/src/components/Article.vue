<template>
  <div class="article">
      <p>{{article}}</p>
  </div>
</template>

<script>
export default {
  name: 'article',
  data() {
    let content = `事务的隔离级别：读未提交；读已提交；可重复读；串行；自我逃避，自我惩罚；实例-文档-源码`
    return {
      article: content
    }
  }
}
</script>

<style scoped>
p {
  align-self: auto;
}
</style>
