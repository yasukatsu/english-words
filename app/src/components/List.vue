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
table, td, th {
  border: 1px #333 solid;
}
thead {
  background: #fda756;
}
tbody {
  background: #fae5d1;
}
.list {
  margin: 0px 60px 60px 60px;
}
</style>
