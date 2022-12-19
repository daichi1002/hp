import { Language } from "~~/types/language";
import fs from "fs";

export const getLanguages = async () => {
  const config = useRuntimeConfig();
  const { data, error } = await useFetch<Language>(
    config.public.PUBLIC_BACKEND_URL + "languages"
  );

  if (!data.value || error.value) {
    console.log(error.value?.message);
  }

  fs.writeFile(`language.json`, JSON.stringify(data.value), (err) => {
    if (err) {
      console.log("JSONファイル書き込みに失敗しました");
    }
  });
};
