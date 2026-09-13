const annotationLabels = {
  de: "Anmerkung {number}",
  en: "Annotation {number}",
  es: "Anotación {number}",
  fa: "یادداشت {number}",
  fr: "Annotation {number}",
  ru: "Примечание {number}",
  zh: "注释 {number}",
};

document$.subscribe(() => {
  const language = document.documentElement.lang.split("-")[0];
  const label = annotationLabels[language] || annotationLabels.en;

  document.querySelectorAll(".md-annotation__index").forEach((link) => {
    const number = link.hash.match(/_annotation_(\d+)$/)?.[1];
    if (number) {
      link.setAttribute("aria-label", label.replace("{number}", number));
    }
  });

  document.querySelectorAll("a.glightbox[data-title]").forEach((link) => {
    if (!link.hasAttribute("aria-label")) {
      link.setAttribute("aria-label", link.dataset.title);
    }
  });
});
