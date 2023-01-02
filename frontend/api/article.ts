import { Article } from "~~/types/article";

export const getArticle = async (id: string | string[]) => {
  const config = useRuntimeConfig();
  const { data, error } = await useFetch<Article>(
    config.public.PUBLIC_BACKEND_URL + `article/${id}`
  );

  if (!data.value || error.value) {
    console.log(error.value?.message);
  }

  return data;
};
