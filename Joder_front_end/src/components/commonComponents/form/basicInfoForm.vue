<template>
  <v-text-field
    v-model="name.value.value"
    :counter="10"
    :error-messages="name.errorMessage.value"
    label="Name"
    variant="underlined"
  ></v-text-field>

  <v-text-field
    v-model="email.value.value"
    :error-messages="email.errorMessage.value"
    label="E-mail"
    variant="underlined"
  ></v-text-field>

  <v-text-field
    v-model="birthday.value.value"
    :error-messages="birthday.errorMessage.value"
    label="生日 格式: 1993-01-01"
    variant="underlined"
  ></v-text-field>

  <v-select
    v-model="userType.value.value"
    :items="seeking"
    item-title="caractor"
    item-value="value"
    :rules="[(value: number) => value != -1 || '請選擇您的角色']"
    label="我的角色"
    variant="underlined"
    required
  ></v-select>

  <v-select
    v-model="gender.value.value"
    :items="genderIdentity"
    item-title="gender"
    item-value="value"
    :rules="[(value: number) => value != -1 || '請選擇您的性別']"
    label="性別"
    variant="underlined"
    required
  ></v-select>
  <v-card-actions>
    <v-btn elevation="4" density="comfortable" rounded="lg" v-if="showPrev" @click="callDecrease">
      上一頁
    </v-btn>
    <v-spacer></v-spacer>
    <v-btn elevation="4" density="comfortable" rounded="lg" v-if="showNext" @click="nextPage">
      下一頁
    </v-btn>
  </v-card-actions>
</template>

<script lang="ts">
import { useFormStore } from '@/stores/formStore'
import { defineComponent} from 'vue'
import { useField, useForm } from 'vee-validate'

export default defineComponent({
  name: 'BasicInfoForm',
  props: ['showPrev', 'showNext'],
  emits: ['callDecrease', 'callIncrease'],
  setup(props, context) {
    const formStore = useFormStore()
    const callIncrease = (): void => {
      context.emit('callIncrease')
    }
    const callDecrease = (): void => {
      context.emit('callDecrease')
    }
    const seeking = [
      { caractor: '求職者', value: 1 },
      { caractor: '獵頭', value: 2 },
      { caractor: 'HR', value: 3 }
    ]

    const genderIdentity = [
      { gender: '男性', value: 1 },
      { gender: '女性', value: 2 },
      { gender: '其他性別認同', value: 3 },
      { gender: '略過', value: 0 }
    ]

    const { handleSubmit } = useForm({
      validationSchema: {
        name(value: string) {
          if (value?.length >= 2) return true

          return 'Name needs to be at least 2 characters.'
        },
        email(value: string) {
          if (
            /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/i.test(
              value
            )
          )
            return true

          return '不是正確的信箱格式'
        },
        userType(value: number) {
          if (typeof value == 'number') {
            return true
          }
          return '選取角色'
        }
      }
    })

    const name = useField('name')
    if (formStore.form.name) {
      name.value.value = formStore.form.name
    }

    const email = useField('email')
    if (formStore.form.email) {
      email.value.value = formStore.form.email
    }

    const userType = useField('userType')
    if (formStore.form.userType) {
      userType.value.value = formStore.form.userType
    }

    const birthday = useField('birthday')
    if (formStore.form.birthday) {
      birthday.value.value = formStore.form.birthday
    }

    const gender = useField('gender')
    if (formStore.form.birthday) {
      gender.value.value = formStore.form.gender
    }

    const nextPage = handleSubmit(() => {
      formStore.form.name = name.value.value
      formStore.form.email = email.value.value
      formStore.form.userType = userType.value.value
      formStore.form.birthday = birthday.value.value
      formStore.form.gender = gender.value.value

      callIncrease()
    })

    return {
      seeking,
      userType,
      email,
      name,
      gender,
      genderIdentity,
      birthday,
      formStore,
      nextPage,
      callDecrease
    }
  }
})
</script>

<style></style>
