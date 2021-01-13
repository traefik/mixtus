const logger = require("./logger");
const { MAX_TITLE_LENGTH, ROOT_URL, titleCase } = require("./util");

/**
 * Adds a Suffix in a format | product-name | version to
 * each page title.
 * @param {HTMLElement} titleElement
 * @param {string} productName
 * @param {string} version
 * @param {string} filePath
 */
exports.addTitleSuffix = (titleElement, productName, version, filePath) => {
  const productNameTitleCase = titleCase(productName.replace('-', ' '));
  const suffix = ` | ${productNameTitleCase} | ${version}`;
  const titleContent = titleElement.innerHTML;

  if (!titleContent.includes(suffix)) {
    let newTitle = `${titleContent.replace(` - ${productNameTitleCase}`, '')}${suffix}`;

    if (newTitle.length > MAX_TITLE_LENGTH) {
      const maxNewTitleLength = MAX_TITLE_LENGTH - suffix.length;
  
      newTitle = `${titleContent.substr(0, maxNewTitleLength - 3).trim()}...${suffix}`;
    }
  
    logger.info(`[title] ${filePath} Adding title suffix. New title: ${newTitle}`);
    titleElement.innerHTML = newTitle;
  }
}

/**
 * Adds to a document a link tag pointing to the canonical page.
 * @param {HTMLDocument} document
 * @param {string} productName
 * @param {string} filePath
 */
exports.addLinkCanonicalURL = (document, productName, filePath) => {
  const link = document.createElement('link');
  link.rel = 'canonical';
  link.href = `${ROOT_URL}/${productName}/`;

  if (!document.head.querySelector('link[rel="canonical"]')) {
    logger.info(`[canonical] ${filePath} Adding canonical link`);
    document.head.appendChild(link);
  }
}

/**
 * Adds to a document a meta robots with index and nofollow.
 * @param {HTMLDocument} document
 * @param {string} filePath
 */
exports.addMetaNofollow = (document, filePath) => {
  const meta = document.createElement('meta');
  meta.name = 'robots';
  meta.content = 'index, nofollow';

  if (!document.head.querySelector('meta[name="robots"][content="index, nofollow"]')) {
    logger.info(`[robots] ${filePath} Adding meta robots`);
    document.head.appendChild(meta);
  }
}

/**
 * Adds to a document a meta description if available as a hidden input.
 * @param {HTMLDocument} document
 * @param {string} filePath
 */
exports.addMetaDescription = (document, filePath) => {
  const content = document.querySelector('#meta-description');

  if (content) {
    const meta = document.createElement('meta');
    meta.name = 'description';
    meta.content = content.value;
  
    if (document.head.querySelector('meta[name="description"]')) {
      logger.info(`[description] ${filePath} Updating meta description`);
      document.head.querySelector('meta[name="description"]').content = content.value;
    } else {
      logger.info(`[description] ${filePath} Adding meta description`);
      document.head.appendChild(meta);
    }
  }
}
