<template>
  <div class="signin">
    <SigninLogin></SigninLogin>
  </div>
</template>

<script lang="ts" setup>
  import { ref } from "vue";
  import { useRouter } from "vue-router";
  const router = useRouter();
  const showPrompt = ref(false);
  function clickHandler1() {
    // Your existing clickHandler1 code
  }

  function clickHandler2() {
    router.push("/signup");
  }

  definePageMeta({
    layout: false,
  });
</script>

<style>
  .signin {
    height: 30%;
    width: 50%;
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    margin: auto;
  }
</style>
