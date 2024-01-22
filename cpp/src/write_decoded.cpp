#include "main.hpp"

void write_decoded(const std::string &file_name, const std::string &content)
{
    std::ofstream output_file(file_name);

    if (!output_file.is_open())
    {
        std::cerr << "Error opening file for writing: " << file_name << std::endl;
        exit(1);
    }

    output_file << content;

    output_file.close();
}