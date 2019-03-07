<template>
  <d3-manul :nodes="nodesinit" :edges="edgesinit" v-if="resetd3"></d3-manul>
</template>

<script>
import D3Manul from '@/components/D3Pic/manul'

const nodesinit = [
  { 'group': 'a', 'id': 50, 'labels': ['Person'], 'props': { 'born': 1956, 'name': 'Tom Hanks' }, 'type': 'node' },
  { 'group': 'b', 'id': 163, 'labels': ['Movie'], 'props': { 'released': 1995, 'tagline': 'Houston, we have a problem.', 'title': 'Apollo 13' }, 'type': 'node' },
  { 'group': 'm', 'id': 134, 'labels': ['Person'], 'props': { 'born': 1954, 'name': 'Ron Howard' }, 'type': 'node' },
  { 'group': 'a', 'id': 50, 'labels': ['Person'], 'props': { 'born': 1956, 'name': 'Tom Hanks' }, 'type': 'node' },
  { 'group': 'b', 'id': 181, 'labels': ['Movie'], 'props': { 'released': 1992, 'tagline': 'Once in a lifetime you get a chance to do something different.', 'title': 'A League of Their Own' }, 'type': 'node' },
  { 'group': 'm', 'id': 185, 'labels': ['Person'], 'props': { 'born': 1943, 'name': 'Penny Marshall' }, 'type': 'node' },
  { 'group': 'a', 'id': 50, 'labels': ['Person'], 'props': { 'born': 1956, 'name': 'Tom Hanks' }, 'type': 'node' },
  { 'group': 'b', 'id': 57, 'labels': ['Movie'], 'props': { 'released': 1990, 'tagline': 'A story of love, lava and burning desire.', 'title': 'Joe Versus the Volcano' }, 'type': 'node' }]

const edgesinit = [
  { 'name': 'r', 'props': { 'roles': ['Jim Lovell'] }, 'relation': 'ACTED_IN', 'source': 7, 'target': 1, 'value': 1 },
  { 'name': 'x', 'props': {}, 'relation': 'DIRECTED', 'source': 6, 'target': 1, 'value': 1 },
  { 'name': 'r', 'props': { 'roles': ['Jimmy Dugan'] }, 'relation': 'ACTED_IN', 'source': 7, 'target': 4, 'value': 1 },
  { 'name': 'x', 'props': {}, 'relation': 'DIRECTED', 'source': 5, 'target': 4, 'value': 1 },
  { 'name': 'r', 'props': { 'roles': ['Joe Banks'] }, 'relation': 'ACTED_IN', 'source': 7, 'target': 2, 'value': 1 },
  { 'name': 'x', 'props': {}, 'relation': 'DIRECTED', 'source': 4, 'target': 7, 'value': 3 },
  { 'name': 'r', 'props': { 'roles': ['Mr. White'] }, 'relation': 'ACTED_IN', 'source': 7, 'target': 4, 'value': 1 },
  { 'name': 'x', 'props': {}, 'relation': 'DIRECTED', 'source': 3, 'target': 6, 'value': 1 }]

export default {
  components: {
    D3Manul
  },
  data () {
    return {
      resetd3: true,
      nodesinit,
      edgesinit
    }
  },
  methods: {
    refreshd3 () {
      this.resetd3 = false
      this.$nextTick(() => {
        this.resetd3 = true
      })
    }
  }
}
</script>
