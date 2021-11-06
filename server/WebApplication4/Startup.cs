using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.HttpsPolicy;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using WebApplication4.Services;
using WebApplication4.Data;
using Microsoft.EntityFrameworkCore;

namespace WebApplication4
{
    public class Startup
    {
        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        // This method gets called by the runtime. Use this method to add services to the container.
        public void ConfigureServices(IServiceCollection services)
        {
            string server = Configuration["DbServer"] ?? "10.59.1.45";
            string port = Configuration["DbPort"] ?? "1433";
            string user = Configuration["DbUser"] ?? "dndUser";
            string password = Configuration["Password"] ?? "cB345678";
            string database = Configuration["Database"] ?? "dndDb";

            string conString = $"Server={server}, {port};Initial Catalog={database};User ID={user};Password={password}";


            services.AddDbContext<Context>(options =>
            {
                //options.UseSqlServer(Configuration.GetConnectionString("DefaultConnection"));
                //options.UseSqlServer("Server=db;Database=master;User=sa;Password=132;");
                options.UseSqlServer(conString);
            });
            //services.AddRazorPages();
        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            DatabaseManagementService.MigrationInitialisation(app);
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
            }
            else
            {
                app.UseExceptionHandler("/Error");
                // The default HSTS value is 30 days. You may want to change this for production scenarios, see https://aka.ms/aspnetcore-hsts.
                app.UseHsts();
            }

            app.UseHttpsRedirection();
            app.UseStaticFiles();

            app.UseRouting();

            //app.UseAuthorization();
            /*
            app.UseEndpoints(endpoints =>
            {
                endpoints.MapRazorPages();
            });
            */
        }
    }
}
