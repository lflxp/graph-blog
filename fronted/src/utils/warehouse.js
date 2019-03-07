import http from 'axios'

// const ip = 'http://127.0.0.1:8080'
// const ip = 'http://10.128.25.30:8080'

const ip = 'https://api-ops.cloudwalk.work'
const setUrl = ip + '/portal-api/etcd/etcdv3/set'
const getUrl = ip + '/portal-api/etcd/etcdv3/get'
const deleteUrl = ip + '/portal-api/etcd/etcdv3/delete'

export default {
  setInfo (args) {
    return http.put(setUrl, args)
  },
  getInfo (args) {
    return http.post(getUrl, args)
  },
  deleteInfo (deleteKey) {
    // https://github.com/axios/axios/issues/897#issuecomment-343715381
    return http.delete(deleteUrl, { data: { key: deleteKey, recursive: false } })
  },
  delete (args) {
    return http.delete(deleteUrl, { data: args })
  }
}
