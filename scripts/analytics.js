import Script from "next/script"

const Analytics = () => {
  return (
    <Script
      dangerouslySetInnerHTML={{
        __html: `(function(c,l,a,r,i,t,y){
        c[a]=c[a]||function(){(c[a].q=c[a].q||[]).push(arguments)};
        t=l.createElement(r);t.async=1;t.src="https://www.clarity.ms/tag/"+i;
        y=l.getElementsByTagName(r)[0];y.parentNode.insertBefore(t,y);
    })(window, document, "clarity", "script", "clln5ndq7m");`,
      }}
      strategy="afterInteractive"
      type="text/javascript"
    />
  )
}

export default Analytics
