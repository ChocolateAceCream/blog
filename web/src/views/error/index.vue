<template>
  <div class="wrapper">
    <div class="inner">
      <div class="text-wrapper">
        <p style="text-align: left;font-weight:bold;font-size:35px;">404</p>

        <p style="text-align: left;">Current page is missing</p>
        <el-button
          type="primary"
          @click="returnHome"
        >Home</el-button>
        <div>
          <el-link
            style=" font-weight:bold;font-size:15px;"
            type="primary"
            @click="goBack"
          >or, return to previous page</el-link>
        </div>

      </div>
      <img src="/img/404.png">
    </div>
  </div>
</template>

<script>
import router from '@/router'
import { defineComponent } from 'vue'
export default defineComponent({
  name: 'Error',
  setup(props, ctx) {
    const returnHome = () => {
      router.push({name: 'home'})
    }
    const goBack = () => {
      router.go(-1)
    }
    return {
      returnHome,
      goBack
    }
  }

})


</script>

<style lang="scss">
.wrapper{
  height: 100%;
  width: 100%;
  display:flex;
  .text-wrapper{
    width:300px;

  }
  img {
    width:22  0px;
  }
}
.inner{
    margin:auto;
    display: flex;
}
.inner p {
    text-align: center;
    font-size: 24px;
}
.inner .leftPic{
    width: 60px;
    height: 60px;
    margin-left: 44%;
    margin-top: 20px;
}
</style>
