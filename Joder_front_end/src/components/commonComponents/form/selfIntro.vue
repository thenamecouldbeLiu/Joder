<template>
  <v-textarea
    v-model.lazy="formStore.form.selfIntro"
    counter
    label="請輸入簡介"
    maxlength="500"
    single-line
  ></v-textarea>
  <v-card-actions>
    <v-btn elevation="4" density="comfortable" rounded="lg" v-if="showPrev" @click="callDecrease">
      上一頁
    </v-btn>
    <v-spacer></v-spacer>
    <v-btn elevation="4" density="comfortable" rounded="lg" v-if="showNext" @click="callIncrease">
      下一頁
    </v-btn>
  </v-card-actions>
</template>

<script lang="ts">
import { useFormStore } from '@/stores/formStore'
import { defineComponent } from 'vue'
export default defineComponent({
  name: 'SelfIntro',
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

    return { formStore, callDecrease, callIncrease }
  }
})
</script>
