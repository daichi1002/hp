import { Language } from "~~/types/language";
import { Commit } from "~~/types/commit";
import { Ref } from "nuxt/dist/app/compat/capi";

export const useContributionStore = () => {
  const languages = useState<Language[]>("language_state", () => []);
  const commits = useState<Commit[]>("commit_state", () => []);
  return {
    languages: readonly(languages),
    commits: readonly(commits),
    getLanguages: getLanguages(languages),
    getCommits: getCommits(commits),
  };
};

const getLanguages = async (state: Ref<Language[]>) => {
  const config = useRuntimeConfig();
  const { data, error } = await useFetch<Language[]>(
    config.public.PUBLIC_BACKEND_URL + "languages"
  );

  if (!data.value || error.value) {
    return console.log(error.value?.message);
  }

  const orderLang = data.value?.sort((a, b) => b.ratio - a.ratio);

  orderLang.map((e) => {
    state.value.push(e);
  });
};

export const getCommits = async (state: Ref<Commit[]>) => {
  const config = useRuntimeConfig();
  const { data, error } = await useFetch<Commit[]>(
    config.public.PUBLIC_BACKEND_URL + "contributions"
  );

  if (!data.value || error.value) {
    return console.log(error.value?.message);
  }

  const orderCommit = data.value?.sort((a, b) => (b.date < a.date ? 1 : -1));

  orderCommit.map((e) => {
    state.value.push(e);
  });
};
