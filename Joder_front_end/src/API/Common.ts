import type { ApiModel, googleAuthUrl } from '@/beans/apiResponse'
import type { MainUser } from '@/beans/userInfo'
import axios from 'axios'
import { type AxiosInstance, type AxiosResponse } from 'axios'
const instance: AxiosInstance = axios.create({
  baseURL: 'http://localhost:9090/api',
  headers: { 'Content-Type': 'application/json' },
  timeout: 20000
})

export function getGoogleOauthUrl(): Promise<AxiosResponse<ApiModel<googleAuthUrl>>> {
  //should be object<string> but too lazy to wrap it up
  const url: Promise<AxiosResponse<ApiModel<googleAuthUrl>>> = instance.get('/ouath/google/url')
  console.log(url)
  return url
}
export function getUserInfo(
  userId: string | null,
  jwt: string | null
): Promise<AxiosResponse<MainUser>> {
  const res: Promise<AxiosResponse<MainUser>> = instance.get('/getUserInfo', {
    params: { userId: userId, jwt: jwt }
  })
  return res
}
