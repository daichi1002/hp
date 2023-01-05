import { Article } from "~~/types/article";

export const getArticle = async (id: string | string[]) => {
  const config = useRuntimeConfig();
  const url = config.public.PUBLIC_BACKEND_URL + `article/${id}`;

  const res = await $fetch<Article>(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    retry: 3,
  });

  return res;
  // TODO：上記どちらかを採用
  // const { data, error } = await useFetch<Article>(
  //   config.public.PUBLIC_BACKEND_URL + `article/${id}`
  // );

  // if (!data.value || error.value) {
  //   console.log(error.value?.message);
  // }

  // return data;
};

export const listArticles = async () => {
  const config = useRuntimeConfig();
  const url = config.public.PUBLIC_BACKEND_URL + "articles";

  const res = await $fetch<Article[]>(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    retry: 3,
  });

  return res;
  // TODO：上記どちらかを採用
  // const config = useRuntimeConfig();
  // const { data, error } = await useFetch<Article[]>(
  //   config.public.PUBLIC_BACKEND_URL + "articles"
  // );

  // if (!data.value || error.value) {
  //   console.log(error.value?.message);
  // }

  // return data;
};
