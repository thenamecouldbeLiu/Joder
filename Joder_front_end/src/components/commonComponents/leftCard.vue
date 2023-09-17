<template>

  <v-card class="flexcard"
          height="100%"
          elevation="6"
          rounded>
    <v-item-group>
      <v-container>
        <!-- 多計算餘數避免少掉最前面幾個 -->
        <v-row v-for="n in testUsers.length + ((testUsers.length%3 === 0)?0:1)"
               :key="n">
          <v-col v-for="i in testUsers[n-1].length"
                 :key="i"
                 cols="4">
            <v-item>
              <v-card :class="['d-flex align-center']"
                      dark
                      height="100px"
                      width="100px">

                <v-img :src="testUsers[n-1][i-1].picUrl"
                       class="align-end"
                       gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)"
                       height="100%"
                       cover>
                  <v-card-subtitle class="text-white font-weight-black">{{ testUsers[n-1][i-1].name }}</v-card-subtitle>
                </v-img>
              </v-card>
            </v-item>
          </v-col>
        </v-row>
      </v-container>

    </v-item-group>

  </v-card>

</template>

<script lang="ts">
import type { User } from '@/beans/userInfo'
import { useLoginStore } from '@/stores/loginStore'
import { defineComponent } from 'vue'
import { ref, type Ref } from 'vue'

export default defineComponent({
  setup() {
    const loginStore = useLoginStore()
    const otherUser = loginStore.testOtherUser

    const testUsers: Ref<User[][]> = ref([
      [otherUser, otherUser, otherUser],
      [otherUser, otherUser, otherUser],
      [otherUser, otherUser, otherUser],
      [otherUser, otherUser, otherUser],
      [otherUser, otherUser, otherUser],
      [otherUser, otherUser, otherUser],
      [otherUser, otherUser, otherUser],
      [otherUser, otherUser, otherUser],
      [otherUser, otherUser, otherUser]
    ])

    return { testUsers }
  }
})
</script>
<style>
</style>