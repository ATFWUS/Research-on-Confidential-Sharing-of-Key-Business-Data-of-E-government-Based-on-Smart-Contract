import {getStore, removeStore} from '@/controller/utils'
import axios from 'axios'
import router from '@/router/index'
import {Loading} from "element-ui";

axios.defaults.baseURL = '/api';

axios.defaults.withCredentials = true;

export default async (url, data, type = 'GET', isShow) => {
  type = type.toUpperCase();
  // url += ".do";
  let showLoading = !!isShow;
  let loadingInstance
  var options = {
    fullscreen: true,
    background: '#010e2c8a',
    customClass: 'loadingStyle'
  }
 

  //response拦截器`
  axios.interceptors.response.use(
    response => {
      if (response.data.code == 403 || response.data.code == 401 || response.data.code == 10005) {
        router.push('/login')
      }
      if (!response.data.success && !response.data.message) {
        response.data.message = "系统内部错误"
      }
      if (!!loadingInstance) {
        loadingInstance.close()
      }

      return response
    },
    error => {
      return Promise.reject(error)
    },
  )

  switch (type) {
    case "GET":
      return axios.get(url, {
        params: data
      })
      break;
    case "POST":
      return axios.post(url, data, {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}, showLoading)
      break;
    case "POSTJSON":
      return axios.post(url, data, {headers: {'Content-Type': 'application/json'}})
      break;
    case "POSTEXPORT":
      return axios.post(url, data, {
        headers: {'Content-Type': 'application/x-www-form-urlencoded'},
        responseType: 'blob'
      })
      break;
    default:
      // statements_def
      break;
  }

}
