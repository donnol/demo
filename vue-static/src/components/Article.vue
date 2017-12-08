<template>
  <div class="article">
      <p>{{article}}</p>
  </div>
</template>

<script>
export default {
  name: 'article',
  data() {
    let content = `事务的隔离级别：读未提交；读已提交；可重复读；串行；`
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
