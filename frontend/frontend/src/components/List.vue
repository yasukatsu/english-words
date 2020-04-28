<template>
  <div class="list">
    <ul>
      <li v-for="word in words" :key="word.Meanings">{{ word }}</li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'List',
  data () {
    return {
      words: []
    }
  },
  created () {
    axios.get('localhost:8000/list').then(res => {
      this.words = res.data
    })
  }
}
</script>

<style scoped>
h1,
h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
