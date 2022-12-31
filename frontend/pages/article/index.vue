<script setup lang="ts">
const { articles, listArticles } = useArticleStore();
onMounted(() => listArticles);
</script>
<template>
  <div class="container mx-auto my-5 p-5">
    <div class="md:flex no-wrap md:-mx-2">
      <Profile />
      <!-- Right Side -->
      <div class="w-full md:w-9/12 mx-2 h-max">
        <div class="grid grid-cols-2">
          <div v-for="article in articles">
            <div
              class="max-w-sm rounded-lg overflow-hidden border mb-6 ml-2 hover:duration-1000 hover:scale-105"
            >
              <NuxtLink :to="`/article/${article.id}`">
                <div class="px-6 py-4">
                  <div class="font-bold text-xl mb-2">{{ article.title }}</div>
                  <p class="text-gray-700 text-base line-clamp-4">
                    {{ article.content }}
                  </p>
                </div>
                <!-- タグの個数分ループを回す -->
                <div class="flex">
                  <div class="pl-6 pt-4 pb-2" v-for="tag in article.tags">
                    <span
                      class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2 mb-2"
                      >#{{ tag.name }}</span
                    >
                  </div>
                </div>
              </NuxtLink>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
