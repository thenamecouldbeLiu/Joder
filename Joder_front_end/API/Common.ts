import axios from 'axios'
import {type AxiosInstance} from 'axios'
const instance: AxiosInstance =  axios.create(
    {
        baseURL: 'https://your.api.domain.tw/',
        headers: { 'Content-Type': 'application/json' },
        timeout: 20000
      });