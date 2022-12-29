export const calcCareer = () => {
  const today = new Date();
  const startDate = "2021-02-08";

  const formatDate = new Date(startDate);
  const diffDay: number = Math.floor(
    (today.getTime() - formatDate.getTime()) / 86400000
  );
  return diffDay;
};

export const formatDate = (date: Date) => {
  return date.toLocaleDateString();
};

export const formatStringDate = (date: string | null | undefined): string => {
  if (date == null || date == undefined) {
    return "";
  }
  console.log(date);
  const formatDate = new Date(date);
  return formatDate.toISOString().split("T")[0];
};
