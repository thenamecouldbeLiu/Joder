import type { ApiModel } from '@/beans/apiResponse'
import type { MainUser, User } from '@/beans/userInfo'

import { type AxiosResponse } from 'axios'
import { commonPost, instance } from './Common'
import type { paginateRequest } from '@/beans/apiRequest'

export function getUserData(): Promise<AxiosResponse<ApiModel<MainUser>>> {
  const res: Promise<AxiosResponse<ApiModel<MainUser>>> = instance.get('/getUserData')
  return res
}

export function getUnmatched(num: number): Promise<AxiosResponse<ApiModel<any>>> {
  const paginateReq: paginateRequest = { paginateNum: num }
  const res: Promise<AxiosResponse<ApiModel<User>>> = commonPost('/getUnmatched', paginateReq)
  return res
}
