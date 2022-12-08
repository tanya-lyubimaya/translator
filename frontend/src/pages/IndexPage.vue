<template>
  <q-layout view="hHh lpR fFf">
    <q-header reveal elevated class="bg-primary text-white">
      <q-toolbar>
        <q-toolbar-title>
          <q-avatar>
            <img src="https://cdn.quasar.dev/logo-v2/svg/logo-mono-white.svg" />
          </q-avatar>
          Title
        </q-toolbar-title>
      </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view />
      <div class="fit row wrap justify-center items-center content-start">
        <q-select
          v-model="inputLanguage"
          use-input
          input-debounce="0"
          label="Language"
          :options="languages"
          @filter="filterFn"
          style="width: 250px"
          behavior="menu"
        />
        <q-btn
          push
          color="primary"
          round
          icon="autorenew"
          class="offset-1"
          @click="switchLanguage"
        />
        <q-select
          v-model="outputLanguage"
          use-input
          input-debounce="0"
          label="Language"
          :options="languages"
          style="width: 250px"
          behavior="menu"
          class="col-4 offset-1"
        />
      </div>
      <div class="fit row wrap justify-center content-start">
        <q-input
          autogrow
          outlined
          clearable
          use-input
          input-debounce="0"
          color="orange"
          v-model="input"
          type="textarea"
          label="Text"
          class="col-4"
          style="overflow: auto"
          clear-icon="close"
        />
        <q-input
          autogrow
          outlined
          disabled
          color="orange"
          v-model="output"
          type="textarea"
          label="Translation"
          class="col-4 offset-1"
          style="overflow: auto"
        />
      </div>
      <div class="fit row wrap justify-center content-start">
        <q-btn
          push
          color="white"
          text-color="primary"
          label="Translate"
          @click="sendForTranslation"
        />
      </div>
    </q-page-container>
  </q-layout>
</template>

<script>
import { ref } from "vue";

const axios = require("axios");

export default {
  setup() {
    return {
      output: ref(null),
      inputLanguage: ref("ru"),
      outputLanguage: ref("en"),
    };
  },
  data() {
    return { languages: [], input: "", languageOptions: [] };
  },
  created() {
    const path = "http://localhost:8080/get-languages";
    this.$axios
      .get(path)
      .then((res) => {
        this.languages = res.data;
        this.languageOptions = res.data;
        console.log("Languages", res);
      })
      .catch((err) => {
        this.$q.notify({
          position: this.notificationsPos,
          icon: "warning",
          type: "negative",
          multiLine: true,
          message: "An error occurred: " + err,
        });
      });
  },
  methods: {
    sendForTranslation() {
      const url = "http://localhost:8080/get-translation";

      if (this.inputLanguage == this.outputLanguage) {
        this.$q.notify({
          position: this.notificationsPos,
          icon: "warning",
          type: "negative",
          multiLine: true,
          message:
            "You're trying to translate from " +
            this.inputLanguage +
            " to " +
            this.outputLanguage +
            "!",
        });
      }
      this.$axios
        .post(url, {
          source_lang: this.inputLanguage,
          target_lang: this.outputLanguage,
          source_text: this.input,
        })
        .then((res) => {
          this.output = res.data;
        })
        .catch((err) => {
          this.$q.notify({
            position: this.notificationsPos,
            icon: "warning",
            type: "negative",
            multiLine: true,
            message: "An error occurred: " + err,
          });
        });
    },

    switchLanguage() {
      let tmp = this.inputLanguage;
      this.inputLanguage = this.outputLanguage;
      this.outputLanguage = tmp;
    },

    filterFn(val, update) {
      if (val === "") {
        update(() => {
          this.languageOptions.value = this.languages;
        });
        return;
      }

      update(() => {
        const needle = val.toLowerCase();
        this.languageOptions.value = this.languages.filter(
          (v) => v.toLowerCase().indexOf(needle) > -1
        );
      });
    },
  },
};
</script>
