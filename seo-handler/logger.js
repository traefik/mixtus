const { LOG_LEVEL = 'ERROR' } = process.env;

const logLevels = {
  ERROR: 1,
  WARN: 2,
  INFO: 3,
  ALL: 4,
  DEBUG: 5,
}

module.exports = {
  /**
   * Console error with prefix.
   * @param  {...any} args
   */
  error: (...args) => {
    console.error('[SEO Transforms]', ...args);
  },
  /**
   * Console warning with prefix.
   * @param  {...any} args
   */
  warn: (...args) => {
    if (logLevels[LOG_LEVEL] && logLevels[LOG_LEVEL] > 1) {
      console.warn('[SEO Transforms]', ...args);
    }
  },
  /**
   * Console log with prefix.
   * @param  {...any} args
   */
  info: (...args) => {
    if (logLevels[LOG_LEVEL] && logLevels[LOG_LEVEL] > 2) {
      console.log('[SEO Transforms]', ...args);
    }
  },
};
