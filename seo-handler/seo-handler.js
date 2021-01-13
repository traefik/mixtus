const { 
  HTML_UNDER_VERSION_REGEX,
  VERSION_REGEX,
  SITEMAP_UNDER_VERSION_REGEX,
  HTML_FILE_REGEX
} = require('./util');
const { JSDOM } = require('jsdom');
const fs = require('fs');
const path = require('path');
const logger = require('./logger');
const iterateDirectory = require('./iterator');
const { addTitleSuffix, addLinkCanonicalURL, addMetaNofollow, addMetaDescription } = require('./transforms');

const [ docPath ] = process.argv.slice(2);

if (!fs.existsSync(docPath)) {
  return logger.error(`Cannot find path: ${docPath}. Aborting.`);
}

const productName = path.basename(docPath);

iterateDirectory(docPath, async (filePath) => {
  const dom = await JSDOM.fromFile(filePath);
  const { document } = dom.window;
  const [,version] = VERSION_REGEX.exec(filePath);
  const initialHtml = dom.serialize();

  if (document.head) { 
    const titleElement = document.head.querySelector('title');

    addLinkCanonicalURL(document, productName, filePath);
    addMetaNofollow(document, filePath);

    if (titleElement) {
      addTitleSuffix(titleElement, productName, version, filePath);
    }

    const finalHtml = dom.serialize();
    if (initialHtml !== finalHtml) {
      fs.writeFile(filePath, finalHtml, e => e && logger.error(e));
    }
  }
}, { includePattern: HTML_UNDER_VERSION_REGEX });

iterateDirectory(docPath, async (filePath) => {
  const dom = await JSDOM.fromFile(filePath);
  const { document } = dom.window;
  const initialHtml = dom.serialize();

  addMetaDescription(document, filePath);

  const finalHtml = dom.serialize();
  if (initialHtml !== finalHtml) {
    fs.writeFile(filePath, finalHtml, e => e && logger.error(e));
  }
}, {
  includePattern: HTML_FILE_REGEX,
  excludePattern: VERSION_REGEX,
});

iterateDirectory(docPath, async (filePath) => {
  logger.info(`[sitemap] ${filePath} deleted.`)
  fs.unlink(filePath, e => e && logger.error(e));
}, { includePattern: SITEMAP_UNDER_VERSION_REGEX });
