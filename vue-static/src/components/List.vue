<template>
    <div class="list">
        <h1>1124</h1>
        <ul>
            <li>{{msg}}</li><br>
        </ul>
        <h2><a>{{author}}</a></h2>
    </div>
</template>

<script>
  export default {
    name: 'List',
    data () {
      return {
        msg: '以往寒风躲闪行，今秋尽是路人多。',
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