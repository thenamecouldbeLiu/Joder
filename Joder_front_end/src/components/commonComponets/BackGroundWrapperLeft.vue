

<template>
  <div class="row">
    <div class="container-fluid">
      <nav class="navbar navbar-expand-lg navbar-light bg-light">
        <div class="container-fluid">
          <a class="navbar-brand" href="#">
            <img
              src="@/../public/favicon.ico"
              alt="icon"
              width="30"
              height="24"
              class="d-inline-block align-text-top"
            />
            {{ isLogin }}
          </a>

          <button
            class="navbar-toggler"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#navbarNav"
            aria-controls="navbarNav"
            aria-expanded="false"
            aria-label="Toggle navigation"
          >
            <span class="navbar-toggler-icon"></span>
          </button>
          <div class="collapse navbar-collapse" id="navbarNav">
            <ul class="navbar-nav">
              <li class="nav-item">
                <a class="nav-link active" aria-current="page" href="#">Matched HH</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="#">See Jobs</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="#">picked Jobs</a>
              </li>
            </ul>
          </div>
        </div>
      </nav>
    </div>

    <div class="row">
      <router-view />
      <div class="col-12">One of three columns</div>
    </div>
  </div>
</template>

<script lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { useLoginStore } from '@/stores/loginStore'
import { storeToRefs } from 'pinia'
export default {
  setup() {
    const loginStore = useLoginStore()
    const { userInfo, isLogin } = storeToRefs(loginStore)

    // expose to template and other options API hooks
    return {
      userInfo,
      isLogin
    }
  },

  mounted() {}
}
</script>
  
  <style scoped>
</style>