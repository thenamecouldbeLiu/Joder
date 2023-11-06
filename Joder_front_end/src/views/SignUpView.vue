<template>
  <form-wrapper></form-wrapper>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import formWrapper from '@/components/commonComponents/form/formWrapper.vue'

export default defineComponent({
  name: 'SignupView',
  components: { formWrapper }
})
</script>
