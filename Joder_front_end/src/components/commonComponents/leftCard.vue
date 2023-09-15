<template>
  <v-card class="flexcard"
          height="100%"
          elevation="6"
          rounded>
    <v-item-group>
      <v-container>
        <v-row v-for="n in Math.ceil(testUsers.length/3)+1"
               :key="n"
               cols="4">
          <v-col v-for="i in 3"
                 :key="i">
            <v-item>
              <v-card :class="['d-flex align-center']"
                      dark
                      height="100%"
                      v-if="i-1< testUsers[n-1].length">
                <div class=" text-center">
                  {{  i-1< testUsers[n-1].length }}
                  {{testUsers[n-1][i-1]}}
                </div>
              </v-card>
            </v-item>
          </v-col>
        </v-row>
      </v-container>

    </v-item-group>

  </v-card>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { ref, type Ref } from 'vue'

export default defineComponent({
  setup() {
    const testUsers: Ref<string[][]> = ref([['first', 'second', 'third'], ['4']])

    return { testUsers }
  }
})
</script>