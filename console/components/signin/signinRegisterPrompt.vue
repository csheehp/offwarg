<template>
  <div class="q-pa-md fit row justify-center q-gutter-md;">
    <q-card class="q-pa-md; q-gutter-md; q-ma-md; col-8">
      <q-card-section>
        <div class="text-h6 q-py-md">Account Registration</div>
        <div class="text-h7 q-py-md">Enter your Account name and user data</div>
      </q-card-section>
      <q-card-section>
        <q-input
          color="grey-3"
          label-color="orange"
          class="q-w-50 shadow-2"
          outlined
          v-model="account_name"
          label="Account Name"
        >
        </q-input>
        <div class="row q-mt-md">
          <q-input
            color="grey-3"
            label-color="orange"
            class="q-w-50 shadow-2 col-6"
            outlined
            v-model="account_name"
            label="First Name"
          >
          </q-input>
          <q-input
            color="grey-3"
            label-color="orange"
            class="q-w-50 shadow-2 col-6 q-pl-sm"
            outlined
            v-model="account_name"
            label="Last Name"
          >
          </q-input>
        </div>
        <div class="q-pt-md">
          <q-input
            color="grey-3"
            label-color="orange"
            class="q-w-50 shadow-2"
            outlined
            v-model="account_name"
            label="Email"
          >
          </q-input>
        </div>
      </q-card-section>
      <q-card-actions align="around">
        <q-btn no-caps outline padding="xs lg" color="grey" @click="emitterSignin">Sign In</q-btn>
        <q-btn no-caps padding="xs lg" color="primary">Register</q-btn>
      </q-card-actions>
    </q-card>
  </div>
</template>

<script lang="ts" setup>
  import { ref, defineEmits } from "vue";
  const account_name = ref("");
  const emit = defineEmits(["register"]);

function emitterSignin() {
  emit("register", false);
}
</script>
<style></style>
