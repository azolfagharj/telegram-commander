(function () {
  var RTL_LANGS = {
    ar: true,
    fa: true,
    he: true,
    ur: true,
    ps: true,
    sd: true,
    yi: true,
    ckb: true,
    dv: true,
  };

  var LABELS = {
    fa: { copy: "کپی", copied: "کپی شد" },
    ar: { copy: "نسخ", copied: "تم النسخ" },
    he: { copy: "העתק", copied: "הועתק" },
    ur: { copy: "کاپی", copied: "کاپی ہو گیا" },
    en: { copy: "Copy", copied: "Copied" },
  };

  function pageLang() {
    var lang = (document.documentElement.lang || "").toLowerCase().split("-")[0];
    return lang || "en";
  }

  function pageDir() {
    var html = document.documentElement;
    var explicit = (html.getAttribute("dir") || "").toLowerCase();
    if (explicit === "rtl" || explicit === "ltr") {
      return explicit;
    }
    var computed = window.getComputedStyle(html).direction;
    if (computed === "rtl" || computed === "ltr") {
      return computed;
    }
    return RTL_LANGS[pageLang()] ? "rtl" : "ltr";
  }

  function labels() {
    var lang = pageLang();
    return LABELS[lang] || LABELS.en;
  }

  function codeText(pre) {
    var code = pre.querySelector("code");
    var text = code ? code.textContent : pre.textContent;
    return (text || "").replace(/\n$/, "");
  }

  function copyText(text) {
    if (navigator.clipboard && navigator.clipboard.writeText) {
      return navigator.clipboard.writeText(text);
    }
    return new Promise(function (resolve, reject) {
      var area = document.createElement("textarea");
      area.value = text;
      area.setAttribute("readonly", "");
      area.style.position = "fixed";
      area.style.insetInlineStart = "-9999px";
      document.body.appendChild(area);
      area.select();
      try {
        document.execCommand("copy");
        resolve();
      } catch (err) {
        reject(err);
      } finally {
        document.body.removeChild(area);
      }
    });
  }

  function addButtons() {
    var text = labels();
    var dir = pageDir();

    document.querySelectorAll("pre").forEach(function (pre) {
      if (pre.closest(".az-copy-wrap")) {
        return;
      }

      var wrap = document.createElement("div");
      wrap.className = "az-copy-wrap";
      wrap.dir = dir;
      pre.parentNode.insertBefore(wrap, pre);
      wrap.appendChild(pre);

      var btn = document.createElement("button");
      btn.type = "button";
      btn.className = "az-copy-btn";
      btn.dir = dir;
      btn.textContent = text.copy;
      btn.setAttribute("aria-label", text.copy);
      wrap.appendChild(btn);

      btn.addEventListener("click", function () {
        copyText(codeText(pre)).then(function () {
          btn.classList.add("is-copied");
          btn.textContent = text.copied;
          window.setTimeout(function () {
            btn.classList.remove("is-copied");
            btn.textContent = text.copy;
          }, 1600);
        });
      });
    });
  }

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", addButtons);
  } else {
    addButtons();
  }
})();
