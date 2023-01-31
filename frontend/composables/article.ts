import { Article } from "~~/types/article";
import { Ref } from "nuxt/dist/app/compat/capi";

export const useArticleStore = () => {
  const articles = useState<Article[]>("articles_state", () => []);

  return {
    articles: readonly(articles),
    listArticles: listArticles(articles),
  };
};

const listArticles = async (state: Ref<Article[]>) => {
  const config = useRuntimeConfig();
  const { data, error } = await useFetch<Article[]>(
    config.public.PUBLIC_BACKEND_URL + "articles"
  );

  if (!data.value || error.value) {
    console.log(error.value?.message);
  }

  data.value?.map((e) => {
    state.value.push(e);
  });
};

const getArticle = async (id: string | string[]) => {
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
