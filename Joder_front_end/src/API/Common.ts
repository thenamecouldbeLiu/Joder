import type { ApiModel, googleAuthUrl } from '@/beans/apiResponse'
import axios from 'axios'

import { type AxiosInstance, type AxiosResponse } from 'axios'

export const instance: AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_BASE_URL + '/api',
  headers: { 'Content-Type': 'application/json' },
  timeout: 20000,
  withCredentials: true
})
export function getGoogleOauthUrl(): Promise<AxiosResponse<ApiModel<googleAuthUrl>>> {
  //should be object<string> but too lazy to wrap it up
  const url: Promise<AxiosResponse<ApiModel<googleAuthUrl>>> = instance.get('/ouath/google/url')
  return url
}

export function commonPost(url: string, req: object): Promise<AxiosResponse<ApiModel<any>>> {
  const res: Promise<AxiosResponse<ApiModel<any>>> = instance.post(url, req)
  return res
}
