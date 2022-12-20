import { Language } from "~~/types/language";

export const getLanguageData = async () => {
  const config = useRuntimeConfig();
  const { error } = await useFetch(
    config.public.PUBLIC_BACKEND_URL + "commit_data"
  );

  if (error.value) {
    console.log(error.value?.message);
  }
};

export const getLanguages = async () => {
  const config = useRuntimeConfig();
  const { data, error } = await useFetch<Language[]>(
    config.public.PUBLIC_BACKEND_URL + "languages"
  );

  if (!data.value || error.value) {
    console.log(error.value?.message);
  }

  return data.value?.sort((a, b) => b.ratio - a.ratio);
};
