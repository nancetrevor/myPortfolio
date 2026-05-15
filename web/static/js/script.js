const yearEl = document.getElementById("year");
if (yearEl) {
  yearEl.textContent = new Date().getFullYear();
}

const revealElements = document.querySelectorAll(".reveal");
const sectionLinks = document.querySelectorAll("[data-section-link]");
const sections = [...document.querySelectorAll("main .section[id]")];

if ("IntersectionObserver" in window && revealElements.length > 0) {
  const observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          entry.target.classList.add("visible");
          observer.unobserve(entry.target);
        }
      });
    },
    { threshold: 0.15 }
  );

  revealElements.forEach((el) => {
    if (el.classList.contains("intro-rail")) {
      el.classList.add("visible");
      return;
    }

    observer.observe(el);
  });
} else {
  revealElements.forEach((el) => el.classList.add("visible"));
}

if ("IntersectionObserver" in window && sectionLinks.length > 0 && sections.length > 0) {
  const linkMap = new Map(
    [...sectionLinks].map((link) => [link.getAttribute("href"), link])
  );

  const sectionObserver = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (!entry.isIntersecting) {
          return;
        }

        sectionLinks.forEach((link) => {
          link.classList.remove("active");
          link.removeAttribute("aria-current");
        });

        const activeLink = linkMap.get(`#${entry.target.id}`);
        if (activeLink) {
          activeLink.classList.add("active");
          activeLink.setAttribute("aria-current", "true");
        }
      });
    },
    {
      rootMargin: "-30% 0px -45% 0px",
      threshold: 0.1,
    }
  );

  sections.forEach((section) => sectionObserver.observe(section));
}
