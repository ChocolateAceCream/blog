<script>
import { getUserList } from '@/api/auth'
import { ref, onMounted, toRefs, defineComponent, computed, reactive } from 'vue'
import { trimStart, partition } from 'lodash-es'
import { sessionStore } from '@/store/sessionStore'

export default defineComponent({
  props: {
    wifiItem: {
      type: Object,
      default: {},
    },
  },
  setup(props, ctx) {
    const store = sessionStore()
    const { name, counter } = store
    const state = reactive({
      lodashDemo: partition([1, 2, 3, 4], n => n % 2),

      doubleCount: computed(() => {
        return store.doubleCount
      }),
    })
    var array = [1, 2, 3]
    console.log(trimStart('-_-abc-_-', '_-', '1'))

    onMounted(() => {
      userList()
    })

    const addCounter = () => {
      console.log("counter: ", counter)
      store.increment()
    }

    const userList = async () => {
      const { data: res } = await getUserList({
        params: {
          pageSize: 15,
          pageNumber: 1,
        }
      })
      console.log('------userlist----', res)
      console.log(name)
      console.log("counter: ", counter)
    }


    return {
      addCounter,
      ...toRefs(state),
      store,
    };
  },

})

</script>

<template>

  <div class="card">
    <button
      type="button"
      @click="addCounter"
    >count is {{ store.counter }}</button>
    <p>double count is {{ doubleCount }}</p>
    <p>
      Edit
      <code>components/HelloWorld.vue</code> to test HMR
    </p>
  </div>
  <span>
    {{ $moment("20111031", "YYYYMMDD").fromNow() }}
  </span>
  <p>
    Check out
    <a
      href="https://vuejs.org/guide/quick-start.html#local"
      target="_blank"
    >create-vue</a>, the official Vue + Vite starter
  </p>
  <p>
    Install
    <a
      href="https://github.com/johnsoncodehk/volar"
      target="_blank"
    >Volar</a>
    in your IDE for a better DX
  </p>
  <p class="read-the-docs">Click on the Vite and Vue logos to learn more</p>
</template>

<style scoped>
.read-the-docs {
  color: #888;
}
</style>
