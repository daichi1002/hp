<script setup lang="ts">
import { getArticle } from "~/api/article";
import { formatStringDate } from "~/util/date";
import MdEditor from "md-editor-v3";
import "md-editor-v3/lib/style.css";
const router = useRoute();
const { id } = router.params;
const article = await getArticle(id);
const content = ref(article.content);
</script>
<template>
  <div class="container mx-auto my-5 p-5">
    <div class="md:flex no-wrap md:-mx-2">
      <Profile />
      <!-- Right Side -->
      <div class="w-full md:w-9/12 mx-2 h-max border rounded-lg p-5">
        <div class="text-center text-3xl font-bold">{{ article?.title }}</div>
        <div class="flex justify-center">
          <div class="" v-for="tag in article?.tags">
            <span
              class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 m-5"
              >#{{ tag.name }}</span
            >
          </div>
        </div>
        <div class="text-center text-sm font-semibold text-gray-700">
          {{ formatStringDate(article?.updatedAt) }}
        </div>
        <md-editor v-model="content" preview-only previewTheme="vuepress" />
      </div>
    </div>
  </div>
</template>
