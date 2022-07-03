<!DOCTYPE xsl:stylesheet [
<!-- Namespaces which will be excluded from the result -->
<!ENTITY ix    "http://www.xbrl.org/2008/inlineXBRL">
<!ENTITY xhtml "http://www.w3.org/1999/xhtml">
]>
<xsl:stylesheet version="1.0"
  xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
  xmlns:ix="&ix;"
  xmlns:xbrli="http://www.xbrl.org/2003/instance"
  xmlns:link="http://www.xbrl.org/2003/linkbase"
  xmlns:xlink="http://www.w3.org/1999/xlink"
  xmlns:x="&xhtml;"
  exclude-result-prefixes="ix xbrli link xlink x">

<xsl:template match="/*">
  <TargetDocument default="yes">
    <xsl:call-template name="narrative"/>
  </TargetDocument>
</xsl:template>
<xsl:template name="narrative">
  <xsl:apply-templates mode="completeText" select="."/>
</xsl:template>
<xsl:template mode="completeText" match="ix:exclude"/>
<xsl:template mode="completeText"
            match="ix:nonFraction
                  | ix:nonNumeric
                  | ix:tuple">
  <xsl:value-of select="./child::*" />
</xsl:template>




</xsl:stylesheet>