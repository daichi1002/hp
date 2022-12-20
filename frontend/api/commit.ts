import { Commit } from "~~/types/commit";
export const getCommits = async () => {
  const config = useRuntimeConfig();
  const { data, error } = await useFetch<Commit[]>(
    config.public.PUBLIC_BACKEND_URL + "contributions"
  );

  if (!data.value || error.value) {
    console.log(error.value?.message);
  }

  return data.value?.sort((a, b) => (b.date < a.date ? 1 : -1));
};
