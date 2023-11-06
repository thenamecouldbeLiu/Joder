<template>
  <v-row justify="center">
    <v-dialog v-model="openForm" persistent width="1024">
      <v-card>
        <v-alert
          type="success"
          title="Alert title"
          :text="sucessMsg"
          closable
          close-label="Close Alert"
          v-if="showSucessMsg"
        ></v-alert>
        <v-alert
          type="error"
          title="Alert title"
          :text="errMsg"
          v-if="showErrMsg"
          closable
          close-label="Close Alert"
        ></v-alert>
        <v-card-title>
          <span class="text-h5" aria-disabled="true"
            >看來你還不是會員，填寫基本資料來搜尋潛在的人才與機會吧！</span
          >
        </v-card-title>
        <v-card-text>
          <component
            :is="formComps[idx]"
            :showPrev="idx > 0"
            :showNext="idx < formComps.length - 1"
            @callIncrease="increaseIdx"
            @callDecrease="decreaseIdx"
          ></component>
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            elevation="4"
            density="comfortable"
            rounded="lg"
            size="large"
            color="#5865f2"
            variant="flat"
            @click="submit"
            v-if="idx == formComps.length - 1"
          >
            上傳
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>
<script lang="ts">
import type { MainUser, InsertUser } from '@/beans/userInfo'
import { defineComponent, ref, type Ref } from 'vue'
import SelfIntro from './selfIntro.vue'
import basicInfoForm from './basicInfoForm.vue'
import { type AxiosResponse } from 'axios'
import { useFormStore } from '@/stores/formStore'
import { commonPost } from '@/API/Common'
import type { ApiModel } from '@/beans/apiResponse'
import router from '@/router'
import { useLoginStore } from '@/stores/loginStore'

export default defineComponent({
  components: { SelfIntro, basicInfoForm },
  name: 'form-Wrapper',
  setup() {
    let userInfo: InsertUser = {}
    const showSucessMsg: Ref<boolean> = ref(false)
    const showErrMsg: Ref<boolean> = ref(false)
    const sucessMsg: Ref<string> = ref('')
    const errMsg: Ref<string> = ref('')
    const formComps: Ref<any[]> = ref(['basic-info-form', 'self-intro']) //form that need to fill
    const idx: Ref<number> = ref(0)
    const openForm: Ref<boolean> = ref(true) //if the logger need to fill form in order to sign up
    const formStore = useFormStore()
    const increaseIdx = (): void => {
      if (idx.value < formComps.value.length - 1) {
        ++idx.value
      }
    }
    const decreaseIdx = (): void => {
      if (idx.value > 0) {
        --idx.value
      }
    }
    const submit = () => {
      commonPost('/InsertUser', formStore.form)
        .then((res: AxiosResponse<ApiModel<MainUser>>) => {
          console.log(res.data.resultCode)
          if (res.data.resultCode === '0000') {
            showSucessMsg.value = true
            const userStore = useLoginStore()
            userStore.userInfo = {
              name: res.data.resultBody.name,
              userId: res.data.resultBody.userId,
              email: res.data.resultBody.name,
              gender: res.data.resultBody.gender,
              lastLoginTime: res.data.resultBody.lastLoginTime
            }
            router.push({ name: 'LoginPage' })
          } else {
            console.log('err')
            errMsg.value = res.data.resultMessage
            showErrMsg.value = true
          }
        })
        .catch((err) => {
          errMsg.value = '未知錯誤發生，請稍後再試 ' + err
          showErrMsg.value = true
        })
    }

    return {
      formComps,
      idx,
      openForm,
      submit,
      userInfo,
      increaseIdx,
      decreaseIdx,
      showSucessMsg,
      showErrMsg,
      sucessMsg,
      errMsg
    }
  }
})
</script>
