/**
 * URL for the root documentation, used to build URL
 * for canonical meta tags.
 */
const ROOT_URL = "https://doc.traefik.io";

/**
 * Regex that represents a version with slashes to be
 * used in folder names.
 * Ex.: /v2.0/, /v1.10/, /v10.0/
 */
const VERSION_REGEX = /\/(v\d+\.\d+)\//

/**
 * Regex that represents a path that ends with '.html'.
 */
const HTML_FILE_REGEX = /\.html$/;

/**
 * Regex that represents a path that contains a version
 * and ends with '.html'.
 */
const HTML_UNDER_VERSION_REGEX = /\/v\d+\.\d+\/.*\.html$/;

/**
 * Regex that represents a path that contains a version
 * and ends with 'sitemap.xml' or 'sitemap.xml.gz'
 */
const SITEMAP_UNDER_VERSION_REGEX = /\/v\d+\.\d+\/.*sitemap\.xml(.gz)?/;

/**
 * Max allowed length for the page title.
 */
const MAX_TITLE_LENGTH = 65;

/**
 * Capitalize the first letter in a string.
 * @param {string} str 
 */
function capitalize(str) {
  return str.charAt(0).toUpperCase() + str.slice(1);
}

/**
 * Capitalize each first letter in a string.
 * @param {string} titleStr 
 */
function titleCase(titleStr) {
  return titleStr.split(' ').map(capitalize).join(' ');
}

module.exports = {
  ROOT_URL,
  VERSION_REGEX,
  HTML_UNDER_VERSION_REGEX,
  HTML_FILE_REGEX,
  SITEMAP_UNDER_VERSION_REGEX,
  MAX_TITLE_LENGTH,
  titleCase,
  capitalize,
}