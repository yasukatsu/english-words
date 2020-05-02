<template>
  <div class="list">
    <table>
      <thead>
        <tr>
          <th>word</th>
          <th>describe</th>
          <th>pos</th>
          <th>japanese</th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="word in words" :key="word">
          <td>{{word.Spel}}</td>
          <td>{{word.Define}}</td>
          <td>{{word.Pos}}</td>
          <td>{{word.Meaning}}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from 'axios'
var path = 'http://0.0.0.0:8000/list'
export default {
  name: 'List',
  data () {
    return {
      words: []
    }
  },
  created () {
    axios.get(path).then(res => {
      this.words = res.data
    })
    console.log(this.words)
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
