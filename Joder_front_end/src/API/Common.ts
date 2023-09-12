import axios from 'axios'
import {type AxiosInstance, type AxiosResponse} from 'axios'
const instance: AxiosInstance =  axios.create(
    {
        baseURL: 'localhost:9090/api',
        headers: { 'Content-Type': 'application/json' },
        timeout: 20000
      });

export function getGoogleOauthUrl(): Promise<AxiosResponse<any, any>> {
  const url:Promise<AxiosResponse<any, any>> = instance.get("/ouath/google/oauthUrl")
  return url;
}