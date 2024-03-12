export class DateTimeUtil {
  formatToMonthDayYear(date) {
    return this.#format(new Date(date), {
      day: "numeric",
      month: "long",
      year: "numeric",
    });
  }

  /**
   * @param {Date} date
   * @param {Intl.DateTimeFormatOptions} config
   */
  #format(date, config) {
    return new Intl.DateTimeFormat("pt-BR", config).format(date);
  }
}
