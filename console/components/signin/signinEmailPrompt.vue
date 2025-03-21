<template>
  <div class="q-pa-md fit row justify-center q-gutter-md;">
    <q-card class="q-pa-md; q-gutter-md; q-ma-md; col-8">
      <q-card-section>
        <div class="text-h6 q-py-md">Welcome to Warg Studio</div>
        <div class="text-h7 q-py-md">Enter your login information</div>
      </q-card-section>
      <q-card-section>
        <q-input
          color="grey-3"
          label-color="orange"
          class="q-w-50 shadow-2"
          outlined
          v-model="mail"
          label="Email"
        >
          <template v-slot:append>
            <q-icon name="mail" color="orange" />
          </template>
        </q-input>
      </q-card-section>
      <q-card-actions align="around">
        <q-btn
          no-caps
          outline
          padding="xs lg"
          color="grey"
          @click="emitterRegister"
          >Register</q-btn
        >
        <q-btn no-caps :disable="isDisable" padding="xs lg" color="primary"
          >Next</q-btn
        >
      </q-card-actions>
    </q-card>
  </div>
</template>
<script lang="ts" setup>
  import { ref, watch, defineEmits} from "vue";
  const mail = ref("");
  const isDisable = ref(true);

  watch(mail, (newVal) => {
    isDisable.value = !(newVal.length > 1 && newVal.includes("@"));
  });
  const emit = defineEmits(["register"]);

  function emitterRegister() {
    console.log("emitterRegister");
    emit("register", true);
  }
</script>
<style></style>
