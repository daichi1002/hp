import { Languages } from "~~/types/language";

export const getCommitData = async () => {
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
  const { data, error } = await useFetch<Languages>(
    config.public.PUBLIC_BACKEND_URL + "languages"
  );

  if (!data.value || error.value) {
    console.log(error.value?.message);
  }

  return data.value;
};
