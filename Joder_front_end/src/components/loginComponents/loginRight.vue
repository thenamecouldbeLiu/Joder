<template>
  <v-banner icon="$vuetify" lines="one" text="Right Wrapper"> </v-banner>

  <v-container fluid class="scrollable fill-height">
    <v-row>
      <v-layout fill-height>
        <template v-slot: loginButton>
          <v-col cols="12">
            <v-btn
              v-for="item in SSO"
              :key="item.description"
              rounded="xl"
              block
              size="x-large"
              :prepend-icon="item.icon"
              @click="item.method()"
              class="mb-8"
              >{{ item.description }}
            </v-btn>
          </v-col>
        </template>
      </v-layout>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { getGoogleOauthUrl, commonPost } from '../../API/Common'
import type { ApiModel, googleAuthUrl } from '../../beans/apiResponse'
import { type AxiosResponse } from 'axios'
interface SSOInterface {
  description: string
  icon: string
  method: Function
}
export default {
  setup() {
    const facebookLogin = async () => {
      const testPost = { model: 'test' }
      console.log(testPost)
      const testRes: AxiosResponse<Object> = await commonPost('/ouath/google/getUserData', testPost)
      console.log(testRes)
    }

    const googleLogin = async () => {
      const res: AxiosResponse<ApiModel<googleAuthUrl>> = await getGoogleOauthUrl()
      //const userStore = useLoginStore()
      //const { userInfo, isLogin, testUser, testOtherUser } = storeToRefs(userStore)
      const url = res.data.resultBody.url
      console.log(res)
      console.log(url)

      window.location.href = url
    }

    const SSO: SSOInterface[] = [
      {
        description: '使用Goole登入',
        icon: 'mdi-google',
        method: googleLogin
      },
      {
        description: '使用Facebook登入',
        icon: 'mdi-facebook',
        method: facebookLogin
      }
    ]
    return { SSO }
  },

  mounted() {}
}
</script>

<style></style>
